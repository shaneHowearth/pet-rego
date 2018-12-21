#!/bin/bash

# Create an Owner
echo "Create an Owner"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"firstname":"abc","surname":"xyz"}' \
  http://localhost:8000/owner
echo ""

# Create a Pet
echo "Create a Pet"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"abc","species":"xyz", "owner":"1"}' \
  http://localhost:8000/pet
echo ""

# Get an owner
echo "Get Owner 1"
curl http://localhost:8000/owner/1
echo ""

# Get a pet
echo "Get Pet 1"
curl http://localhost:8000/pet/1
echo ""

# Get all owners
echo "Get all owners"
curl http://localhost:8000/owner
echo ""

# Get all pets
echo "Get all pets"
curl http://localhost:8000/pet
echo ""

# Get pets by owner
echo "Get pets by owner"
curl http://localhost:8000/owner/1/pets

echo ""
