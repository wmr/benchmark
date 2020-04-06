import grpc
import cbor
import server_pb2
import server_pb2_grpc

def call(stub):    
    request = server_pb2.DataRequest(name='data')
    #response_iterator = stub.GetData(request)
    response_iterator = stub.GetData(request)
    #print(list(response))
    f = open('resp.data', 'wb')
    for response in response_iterator:
    #print(response)
        #[f.write(it) for it in response.data]
        unpacked = [cbor.loads(it) for it in response.data]
        #print(f"recv from message={len(unpacked)}")
    f.close()

options = [('grpc.max_message_length', 100 * 1024 * 1024)]
with grpc.insecure_channel('[::]:50051', options=options) as channel:   
    stub = server_pb2_grpc.DataServiceStub(channel)

    call(stub)