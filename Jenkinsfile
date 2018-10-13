pipeline {
    agent { docker 'golang:1.11.1'}
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}