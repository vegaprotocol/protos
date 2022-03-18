pipeline {
    agent { label 'general' }
    environment {
        CGO_ENABLED = 0
        GO111MODULE = 'on'
        SLACK_MESSAGE = "Protos CI » <${RUN_DISPLAY_URL}|Jenkins ${BRANCH_NAME} Job>${ env.CHANGE_URL ? " » <${CHANGE_URL}|GitHub PR #${CHANGE_ID}>" : '' }"
    }

    stages {
	stage('unit tests') {
            options { retry(3) }
            steps {
                sh 'go test -v ./... 2>&1 | tee unit-test-results.txt && cat unit-test-results.txt | go-junit-report > vega-unit-test-report.xml'
                junit checksName: 'Unit Tests', testResults: 'vega-unit-test-report.xml'
            }
	}
        stage('Run linters') {
            parallel {
                stage('buf lint') {
                    steps {
                        retry(3) {
                            sh 'buf lint'
                        }
                    }
                }
                stage('proto check') {
                    steps {
                        retry(3) {
                            sh "./script/gettools.sh"
                            sh 'make proto_check'
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
    }
}
