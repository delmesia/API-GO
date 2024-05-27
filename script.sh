#!/bin/bash
echo ""
echo "Performing GET request to localhost:4000/v1/healthcheck..."
curl localhost:4000/v1/healthcheck
echo ""

echo "Performing GET request to localhost:/4000/v1/movies/123"
curl localhost:4000/v1/movies/123
echo ""
