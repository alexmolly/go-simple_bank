#!/bin/bash

ZONE=us-central1-a
NEW_INSTANCE_NAME=gcp-instance-2
DB_NAME=simple_bank


gcloud sql instances create $NEW_INSTANCE_NAME \
--database-version=POSTGRES_14 \
--cpu=1 \
--memory=3840MB \
--zone=$ZONE \
--availability-type=zonal \
--no-backup \
--root-password=secret \
--assign-ip \
--storage-size=10GB \
--storage-type=SSD \
--no-storage-auto-increase \

gcloud sql databases create  $DB_NAME --instance=$NEW_INSTANCE_NAME

