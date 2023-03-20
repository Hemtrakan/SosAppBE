#!groovy

def ENVIRONMENT = 'dev'


def PROJECT_ENV = 'dev'
def PROJECT_DIR1 = 'accounts'
def PROJECT_DIR2 = 'emergency'
def PROJECT_DIR3 = 'hotline'
def PROJECT_DIR4 = 'messenger'


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
                        dir("./${PROJECT_DIR1}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image1}:${BUILD_NUMBER} ."
                        }


                        dir("./${PROJECT_DIR2}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image2}:${BUILD_NUMBER} ."
                        }


                        dir("./${PROJECT_DIR3}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image3}:${BUILD_NUMBER} ."
                        }


                        dir("./${PROJECT_DIR4}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image4}:${BUILD_NUMBER} ."
                        }

//                         def slackImage1 = docker.build("${env.image1}:${BUILD_NUMBER}")
//                         slackImage1.push()
//                         slackImage1.push('latest')
//
//                         def slackImage2 = docker.build("${env.image2}:${BUILD_NUMBER}")
//                         slackImage2.push()
//                         slackImage2.push('latest')
//
//                         def slackImage3 = docker.build("${env.image3}:${BUILD_NUMBER}")
//                         slackImage3.push()
//                         slackImage3.push('latest')
//
//                         def slackImage4 = docker.build("${env.image4}:${BUILD_NUMBER}")
//                         slackImage4.push()
//                         slackImage4.push('latest')

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