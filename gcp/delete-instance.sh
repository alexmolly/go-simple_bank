#!/bin/bash


INSTANCE_NAME=gcp1


echo "deleting instance $INSTANCE_NAME"

gcloud sql instances delete $INSTANCE_NAME
