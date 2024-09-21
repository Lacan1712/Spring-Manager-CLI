pipeline {
    agent { label 'localhost' }

    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }

    stages {
        stage('Clone Repository') {
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: '*/${env.BRANCH_NAME}']], // Clona a branch que está sendo utilizada
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }

        stage('Build for Linux and Windows') {
            steps {
                script {
                    dir("${env.WORKSPACE}") {
                        // Executa o script de build para Linux e Windows
                        sh 'sh scripts/builds/build_develop.sh'
                    }
                }
            }
        }
    }

    post {
        success {
            echo "Build completed successfully."
            archiveArtifacts artifacts: 'build/**', allowEmptyArchive: true
        }
        failure {
            echo "Build failed. Please check the logs."
        }
    }
}
