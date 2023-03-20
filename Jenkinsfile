#!groovy

def ENVIRONMENT = 'dev'

pipeline {
	agent any

	environment {
        image1 = "annabells/sosapp-accounts-service"
        image2 = "annabells/sosapp-emergency-service"
        image3 = "annabells/sosapp-hotline-service"
        image4 = "annabells/sosapp-messenger-service"
        registry = "docker.io"
    }


	stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Print Environment') {
            steps {
                sh('ls -al')
                sh('printenv')
            }
        }

		stage('Build docker image') {
            steps {
                script {
                    docker.withRegistry('', 'dockerhub') {

                        echo("ls -al")
                        echo("BUILD_NUMBER")

                        def slackImage1 = docker.build("${env.image1}:${BUILD_NUMBER}")
                        slackImage.push()
                        slackImage.push('latest')

                        def slackImage2 = docker.build("${env.image2}:${BUILD_NUMBER}")
                        slackImage.push()
                        slackImage.push('latest')

                        def slackImage3 = docker.build("${env.image3}:${BUILD_NUMBER}")
                        slackImage.push()
                        slackImage.push('latest')

                        def slackImage4 = docker.build("${env.image4}:${BUILD_NUMBER}")
                        slackImage.push()
                        slackImage.push('latest')

                        sh('docker logout')
                    }
                }
            }
        }

		stage('Deployment'){
            steps {
                sh "docker-compose up -d"
            }
        }



        stage('tag docker image') {
            steps {
                sh "docker tag ${env.image}:${BUILD_NUMBER} ${env.image1}:latest"
                sh "docker tag ${env.image}:${BUILD_NUMBER} ${env.image2}:latest"
                sh "docker tag ${env.image}:${BUILD_NUMBER} ${env.image3}:latest"
                sh "docker tag ${env.image}:${BUILD_NUMBER} ${env.image4}:latest"
            }
        }

        stage('push docker image') {
            steps {
                h "docker push ${env.image1}:latest"
                h "docker push ${env.image2}:latest"
                h "docker push ${env.image3}:latest"
                h "docker push ${env.image4}:latest"
            }
        }

//         stage('Verify new docker image(s)') {
//             steps {
//                 sh('docker images')
//             }
//         }

	}
}