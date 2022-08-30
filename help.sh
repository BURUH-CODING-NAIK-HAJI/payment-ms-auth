#!/bin/bash

type=$1
name=$2

if [ $type = "controller" ]; then
    foldername=$name"controller"
    newcontrollerpath=app/controller/$foldername/RenameThisFileController.go
    mkdir app/controller/$foldername

    cp app/controller/welcomecontroller/WelcomeController.go $newcontrollerpath

    sed -i -e 's/welcomecontroller/'$foldername'/' $newcontrollerpath
fi  

if [ $type = "service" ]; then
    foldername=$name"service"
    newservicepath=app/service/$foldername/RenameThisFileService.go
    mkdir app/service/$foldername

    cp app/service/welcomeservice/WelcomeService.go $newservicepath

    sed -i -e 's/welcomeservice/'$foldername'/' $newservicepath
fi  

if [ $type = "repository" ]; then
    foldername=$name"repository"
    newrepositorypath=app/repository/$foldername/RenameThisFileRepository.go
    mkdir app/repository/$foldername

    cp app/repository/welcomerepository/WelcomeRepository.go $newrepositorypath

    sed -i -e 's/welcomerepository/'$foldername'/' $newrepositorypath
fi  