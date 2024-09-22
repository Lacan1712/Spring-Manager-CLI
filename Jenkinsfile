pipeline {
    agent { label 'localhost' }

    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }

    stages {
        stage('Clean Workspace') {
            steps {
                cleanWs() // Limpa o workspace
            }
        }

        stage('Clone Repository') {
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: '*/develop']],
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }

        stage('Build for Linux and Windows') {
            steps {
                script {
                    dir("${env.WORKSPACE}") {
                        sh '../scripts/SpringCLI/build_develop.sh'  // Executa o script de build
                    }
                }
            }
        }
    }

    post {
        success {
            echo "Build completed successfully."
        }
        failure {
            echo "Build failed. Please check the logs."
        }
    }
}
