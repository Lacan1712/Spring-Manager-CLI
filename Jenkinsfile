pipeline {
    agent { label 'localhost' }

    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }

    stages {
        stage('Clone Repository') {
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: '*/develop']],  // Use explicitamente a branch 'develop'
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }

        stage('Build for Linux and Windows') {
            steps {
                script {
                    dir("${env.WORKSPACE}") {
                        sh 'sh /scripts/Spring\ CLI/build_develop.sh'  // Executa o script de build
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
