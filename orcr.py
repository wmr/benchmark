#!/usr/local/bin/python3

import pyorc

from concurrent import futures
import logging
import grpc

import server_pb2
import server_pb2_grpc

import pyorc
import cbor2


import cbor2

with open('./data.orc', 'rb') as data:
    reader = pyorc.Reader(data)
    print(reader.schema)
    for row in reader:
        pass #print(row)
