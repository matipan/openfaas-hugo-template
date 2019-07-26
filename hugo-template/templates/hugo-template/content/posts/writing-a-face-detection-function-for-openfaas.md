+++
title = "Writing a face detection function for OpenFaaS"
date = "2019-03-16"
tags = ["go","gocv","openfaas","serverless"]
categories = ["go","serverless"]
+++

There is a new term in town that has been making its way to all of us for a couple of years now: *Serverless*. When I first heard this I was kind of confused as to what it meant. Everybody seemed to have an opinion about it but there were no real answers. The following definition from [serverless-stack](https://serverless-stack.com/chapters/what-is-serverless.html) helped me clarify a bit:

> Serverless computing (or serverless for short), is an execution model where the cloud provider is responsible for executing a piece of code by dynamically allocating the resources. And only charging for the amount of resources used to run the code.

It also claims the following:

> While serverless abstracts the underlying infrastructure away from the developer, servers are still involved in executing our functions.

Essentially we want to write our code without worrying where it will be executed, how it will be scaled and how it will exposed to the rest of the world. [This talk](https://www.youtube.com/watch?v=oNa3xK2GFKY) by [Kelsey Hightower](https://twitter.com/kelseyhightower) does an amazing job in showing off what serverless is all about.

In this post I'll go through how to implement our own Serverless function using OpenFaaS. We are going to create a small program that receives an image or a URL and returns the same image with all the faces marked with a green rectangle. We'll do this using OpenCV and a neural network.

## Prerequisites
For this tutorial I will be using GCP's (Google Cloud Platform) managed Kubernetes service called GKE to deploy OpenFaaS and the face detection program will be built using Go and GoCV. So be sure you meet with the following requirements:

* Go installed on your machine. If you want to test the program locally you'll either need to install GoCV or use one of their available [docker images](https://github.com/denismakogon/gocv-alpine).
* A Kubernetes cluster with kubectl set up. If you want you can use GCP's 300USD free credit. We won't be using much of it. You can follow [this tutorial](https://cloud.google.com/kubernetes-engine/docs/quickstart) to set it up. You can use [Minikube](https://medium.com/@lizrice/getting-started-with-openfaas-on-minikube-8d51987f5bbb) for this if you prefer to test it all locally. I haven't actually followed that blog post so I'm not sure if it's still working.

## OpenFaaS
We already said that while serverless abstracts the underlying infrastructure from your code that doesn't mean there is no infrastructure behind. There are tools that different cloud providers have such as [Google's Cloud Functions](https://cloud.google.com/functions/), [Microsoft's Azure Functions](https://azure.microsoft.com/en-us/services/functions/) and [AWS Lambda](https://aws.amazon.com/lambda/) that allow us to write our serverless functions. But today we are going to be looking at one called [OpenFaaS](https://openfaas.com).

OpenFaaS is a platform that tries to *Make Serverless Functions Simple* (and it does that very well). It can run on top of something like Kubernetes, Docker Swarm, Fargate and others. It makes life easier for developers and operators because of its ease of use and amazing CLI. Check out the [documentation](https://docs.openfaas.com/). It's very extensive and has all the information you'll need.

Before we start writing any code, let's get it up and running with OpenFaaS. Since you already have a Kubernetes cluster, follow [this tutorial](https://github.com/openfaas/faas-netes/blob/master/HELM.md) to deploy it using helm. Once you have it up and running install the [faas-cli](https://github.com/openfaas/faas-netes/blob/master/chart/openfaas/README.md#verify-the-installation) on your machine and [authenticate to your OpenFaaS installation](https://github.com/openfaas/faas-cli#get-started-install-the-cli).

### Functions and templates
All functions in OpenFaaS belong to a [template](https://github.com/openfaas/templates). There are many [templates for different languages and tools](https://github.com/openfaas/templates) already provided. But if none of those meet your needs, it is very easy to create a new one. This is what we'll be doing in this article.

### Writing our function
To perform the face detection we are going to use GoCV and a pre-trained Caffe neural network. This means that our function needs to have OpenCV installed and the model+config files of the neural network. We could use the [dockerfile template](https://github.com/openfaas/templates/tree/master/template/dockerfile) and install our dependencies there, but I like the idea of having a *gocv template* with built-in models for anyone who wants to make use of it.
[Here](https://github.com/matipan/openfaas-gocv-template/) is the template I created in case you want to check it out.

Let's start writing our function! The users will provide the image they want to perform the face detection on as a URL. The function will:

* Do a GET request to that URL and check if it's an image and if the image is of the supported content types (jpg, jpeg and png).
* If it's not then we simply return an error.
* If it is valid we need to:
    1. decode the image,
    2. run a pass through the neural network,
    3. draw rectangles around all of the faces that were found, 
    4. encode the image and return the results.
    
The code to do the detection was extracted from [this example](https://github.com/hybridgroup/gocv/blob/master/cmd/dnn-detection/main.go) of GoCV. I had to make a few minor changes to parse the downloaded image and encode it before returning it, but the important parts are the same.
Download the function's code from [Github](https://github.com/matipan/openfaas-face-finder) and let's deploy it using the faas cli.
Once you've downloaded the function go to the `stack.yml` file and change the following fields:

* `provider > gateway`: set your openfaas gateway URL
* `functions > face-finder > image`: change `matipan` for your own docker hub username.

Now we simply have to deploy the function, this is as simple as running: `faas up`. This command will build the image, push it to the docker registry and deploy it on openfaas. Once the function is deployed head over to the OpenFaaS dashboard and select your face-finder function. Provide a URL([like this one](http://dujye7n3e5wjl.cloudfront.net/photographs/1080-tall/time-100-influential-photos-ellen-degeneres-oscars-selfie-100.jpg)) to an image then select the *Download* option and hit invoke. This should download an image onto your machine, if it all worked it should be something like this:

![Image with all the faces marked by a green rectangle](https://raw.githubusercontent.com/matipan/openfaas-face-finder/master/doc/result.jpg)

Pretty neat, right?

If you are interested in getting into details on how this was implemented you can check out the `function/handler.go` file. It is less than 100 lines of code and is quite straightforward.

If you have any questions feel free to contact me on [Twitter](https://twitter.com/matiaspan26)!
