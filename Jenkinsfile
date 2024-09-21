pipeline {
    agent none
    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }
    stages {
        stage('Preparation') {
            agent { label getAgentForBranch(env.BRANCH_NAME) }
            steps {
                echo "Preparing for build on branch: ${env.BRANCH_NAME}"
            }
        }
        stage('Clone') {
            agent { label getAgentForBranch(env.BRANCH_NAME) }
            steps {
                script {
                    clone(env.BRANCH_NAME)
                }
            }
        }
        stage('Build') {
            steps {
                script {
                    runBuildForBranch(env.BRANCH_NAME)
                }
            }
        }
    }
    post {
        success {
            echo "Builds completed successfully."
        }
        failure {
            echo "Build failed. Please check the logs."
        }
    }
}

def getAgentForBranch(branch) {
    switch (branch) {
        case 'main':
            return 'localhost-main'
        case 'develop':
            return 'localhost'
        case { it.startsWith('feature/') }:
            return 'agent-feature'
        default:
            return 'agent-default'
    }
}

def clone(branch) {
    checkout([$class: 'GitSCM', 
        branches: [[name: '*/' + branch]],
        userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
    ])
}

// Função para o build da branch main
def buildMainBranch() {
    echo 'Building the main branch...'
    node {
        script {
            dir("${env.WORKSPACE}") {  
                sh '''
                go version
                GOOS=linux GOARCH=amd64 go build -o smc
                zip -r smc-linux-amd64 smc
                '''
            }
        }
    }
}

// Função para o build da branch develop
def buildDevelopBranch() {
    echo 'Building the develop branch...'
    node {
        script {
            dir("${env.WORKSPACE}") {  
                sh '''
                go version
                GOOS=linux GOARCH=amd64 go build -o smc
                zip -r smc-linux-amd64 smc
                GOOS=windows GOARCH=amd64 go build -o smc.exe
                zip -r smc-win-amd64 smc.exe
                '''
            }
        }
    }
}

// Função para o build de branches de feature
def buildFeatureBranch(branch) {
    echo "Building feature branch: ${branch}"
    node {
        script {
            dir("${env.WORKSPACE}") {  
                sh '''
                go version
                GOOS=linux GOARCH=amd64 go build -o smc-feature
                zip -r smc-feature-linux-amd64 smc-feature
                '''
            }
        }
    }
}