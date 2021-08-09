pipeline {
    agent { label 'general' }
    environment {
        CGO_ENABLED = 0
        GO111MODULE = 'on'
        SLACK_MESSAGE = "Protos CI » <${RUN_DISPLAY_URL}|Jenkins ${BRANCH_NAME} Job>${ env.CHANGE_URL ? " » <${CHANGE_URL}|GitHub PR #${CHANGE_ID}>" : '' }"
    }

    stages {

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
