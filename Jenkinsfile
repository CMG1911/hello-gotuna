pipeline {
    agent any
    options {
      timestamps ()
      ansiColor('xterm')
    }  
        stages {
          stage('Build') {
            steps {
             echo '\033[42m\033[97mConstrucción\033[0m'
             sh 'docker-compose build'
                }
          }
          stage('Up') {
            steps {
             echo '\033[42m\033[97mLevantamiento\033[0m'
             sh 'docker-compose up -d'
            }
          }
        }
}

