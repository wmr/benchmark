# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import server_pb2 as server__pb2


class DataServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetData = channel.unary_stream(
        '/datasvc.DataService/GetData',
        request_serializer=server__pb2.DataRequest.SerializeToString,
        response_deserializer=server__pb2.DataResponse.FromString,
        )


class DataServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetData(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_DataServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetData': grpc.unary_stream_rpc_method_handler(
          servicer.GetData,
          request_deserializer=server__pb2.DataRequest.FromString,
          response_serializer=server__pb2.DataResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'datasvc.DataService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
