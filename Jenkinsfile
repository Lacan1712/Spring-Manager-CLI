pipeline {
    agent none // Define que não haverá um agent padrão
    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }
    stages {
        stage('Determine Agent') {
            agent {
                label getAgentLabel()
            }
            steps {
                script {
                    echo "Agent selected for branch: ${env.BRANCH_NAME}"
                }
            }
        }
        stage('Clone') {
            agent {
                label getAgentLabel() // Chama a função para obter o label do agent
            }
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: "${env.BRANCH_NAME}"]],
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }
        stage('Build and Zip') {
            agent {
                label getAgentLabel() // Chama a função para obter o label do agent
            }
            steps {
                script {
                    runBranchSpecificScripts()
                }
            }
        }
    }
    post {
        success {
            echo "Build completed successfully for branch ${env.BRANCH_NAME}."
        }
        failure {
            echo "Build failed for branch ${env.BRANCH_NAME}. Please check the logs."
        }
    }
}

// Função para determinar o label do agent com base na branch
def getAgentLabel() {
    switch (env.BRANCH_NAME) {
        case 'develop':
            return 'localhost' // Define o label para a branch develop
        case { it.startsWith('feature/') }:
            return 'feature-agent' // Define o label para branches de feature
        case { it.startsWith('release/') }:
            return 'release-agent' // Define o label para branches de release
        default:
            return 'default-agent' // Define um label padrão para outras branches
    }
}

// Função para executar scripts específicos com base na branch
def runBranchSpecificScripts() {
    switch (env.BRANCH_NAME) {
        case 'develop':
            buildForLinux()
            buildForWindows()
            zipArtifacts()
            break
        case { it.startsWith('feature/') }:
            buildForLinux() // Executa apenas o build para Linux
            break
        case { it.startsWith('release/') }:
            buildForWindows() // Executa apenas o build para Windows
            break
        default:
            echo "No specific build actions defined for branch: ${env.BRANCH_NAME}"
            break
    }
}

// Função para construir para Linux
def buildForLinux() {
    script {
        dir("${env.WORKSPACE}") {
            sh """
                echo "Building for Linux on branch: ${env.BRANCH_NAME}"
                go version  # Verifica se o Go está disponível
                GOOS=linux GOARCH=amd64 go build -o nome-do-app-linux-${env.BRANCH_NAME}
            """
        }
    }
}

// Função para construir para Windows
def buildForWindows() {
    script {
        dir("${env.WORKSPACE}") {
            sh """
                echo "Building for Windows on branch: ${env.BRANCH_NAME}"
                go version  # Verifica se o Go está disponível
                GOOS=windows GOARCH=amd64 go build -o nome-do-app-windows-${env.BRANCH_NAME}.exe
            """
        }
    }
}

// Função para zipar os artefatos
def zipArtifacts() {
    script {
        dir("${env.WORKSPACE}") {
            sh """
                echo "Zipping artifacts..."
                zip nome-do-app-linux-${env.BRANCH_NAME}.zip nome-do-app-linux-${env.BRANCH_NAME}
                zip nome-do-app-windows-${env.BRANCH_NAME}.zip nome-do-app-windows-${env.BRANCH_NAME}.exe
            """
        }
    }
}
