#!/bin/bash

function build {
  echo
  echo "*************************************"
  echo "* BUILDING APPLICATION"
  echo "*************************************"
  echo
  cd src || exit
  go build || exit
  mv app ../
}

build || echo "* BUILD FAILURE"
echo
echo "*************************************"
echo "* FINISHED"
echo "*************************************"
echo