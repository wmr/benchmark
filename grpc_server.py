#!/usr/local/bin/python3
import pyorc
from concurrent import futures
import logging
import grpc

import server_pb2
import server_pb2_grpc


import cbor

class DataServer(server_pb2_grpc.DataServiceServicer):

    def GetData(self, request, context):
        with open('./data.orc', 'rb') as dataf:
            reader = pyorc.Reader(dataf)
            print(reader.schema)

            #data = [cbor.dumps(it) for it in reader]
            chunk = []
            for idx, row in enumerate(reader):
                #chunk.append(cbor.dumps(row))
                chunk.append(row)
                if idx % 10000 == 0:
                    yield server_pb2.DataResponse(data=[cbor.dumps(chunk)])
                    chunk = []
        yield server_pb2.DataResponse(data=[cbor.dumps(chunk)])


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    server_pb2_grpc.add_DataServiceServicer_to_server(DataServer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()