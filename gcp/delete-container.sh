#!/bin/bash


CONTAINER_NAME=gcr.io/my-linux-project-348713/simple_bank


echo "deleting container $CONTAINER_NAME"

gcloud container images delete $CONTAINER_NAME
