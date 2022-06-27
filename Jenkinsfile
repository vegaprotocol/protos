pipeline {
    agent { label 'general' }
    environment {
        CGO_ENABLED = 0
        GO111MODULE = 'on'
        SLACK_MESSAGE = "Protos CI » <${RUN_DISPLAY_URL}|Jenkins ${BRANCH_NAME} Job>${ env.CHANGE_URL ? " » <${CHANGE_URL}|GitHub PR #${CHANGE_ID}>" : '' }"
        GOPATH = "${env.WORKSPACE}/GOPATH"
        GOBIN = "${env.GOPATH}/bin"
        PROTOC_HOME = "${env.WORKSPACE}/PROTOC_HOME"
        PATH = "${env.PROTOC_HOME}/bin:${env.GOBIN}:${env.PATH}"
        PROTOC_VERSION = "3.19.4"
    }

    stages {
        stage('Config') {
            steps {
                sh 'printenv'
                echo "params=${params}"
            }
        }
        stage('unit tests') {
            options { retry(3) }
            steps {
                sh 'go test -v ./... 2>&1 | tee unit-test-results.txt && cat unit-test-results.txt | go-junit-report > vega-unit-test-report.xml'
                junit checksName: 'Unit Tests', testResults: 'vega-unit-test-report.xml'
            }
        }
        stage('Install dependencies') {
            options { retry(3) }
            steps {
                sh label: 'protoc', script: """#!/bin/bash -e
                    PB_REL="https://github.com/protocolbuffers/protobuf/releases"
                    curl -LO \$PB_REL/download/v${env.PROTOC_VERSION}/protoc-${env.PROTOC_VERSION}-linux-x86_64.zip
                    unzip protoc-${env.PROTOC_VERSION}-linux-x86_64.zip -d "${PROTOC_HOME}"
                """
                sh './script/gettools.sh'
                sh 'protoc --version'
                sh 'which protoc'
                sh 'buf --version'
                sh 'which buf'
            }
        }
        stage('Run linters') {
            parallel {
                stage('buf lint') {
                    options { retry(3) }
                    steps {
                        sh 'buf lint'
                    }
                    post {
                        failure {
                            sh 'printenv'
                            echo "params=${params}"
                            sh 'protoc --version'
                            sh 'which protoc'
                            sh 'buf --version'
                            sh 'which buf'
                            sh 'git diff'
                        }
                    }
                }
                stage('proto check') {
                    options { retry(3) }
                    steps {
                        sh 'make proto_check'
                    }
                    post {
                        failure {
                            sh 'printenv'
                            echo "params=${params}"
                            sh 'protoc --version'
                            sh 'which protoc'
                            sh 'buf --version'
                            sh 'which buf'
                            sh 'git diff'
                        }
                    }
                }
            }
        }
    }
    post {
        success {
            retry(3) {
                slackSend(channel: "#tradingcore-notify", color: "good", message: ":white_check_mark: ${SLACK_MESSAGE} (${currentBuild.durationString.minus(' and counting')})")
            }
        }
        failure {
            retry(3) {
                slackSend(channel: "#tradingcore-notify", color: "danger", message: ":red_circle: ${SLACK_MESSAGE} (${currentBuild.durationString.minus(' and counting')})")
            }
        }
        always {
            cleanWs()
        }
    }
}
