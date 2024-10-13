import model
import grpc
from concurrent import futures
# from ..common.api.ModelService_pb2 import *
# from ..common.api.ModelService_pb2_grpc import *
import ultraimport
import configparser


ModelService_pb2 = ultraimport('__dir__/../common/api/ModelService_pb2.py')
ModelService_pb2_grpc = ultraimport(
    '__dir__/../common/api/ModelService_pb2_grpc.py')


class ModelService(ModelService_pb2_grpc.ModelServiceServicer):
    def __init__(self):
        self.model = model.Model()

    def TextEmbedding(self, request, context):
        return TextEmbeddingResponse(embedding=self.model.get_embedded_text(request.text))

    def ImageEmbedding(self, request_iterator, context):
        image_data_chunks = []
        for request in request_iterator:
            image_data_chunks.append(request.chunk)
        full_image_data = b''.join(image_data_chunks)
        return ModelService_pb2.ImageEmbeddingResponse(embedding=self.model.get_embedded_image(full_image_data))


def serve():

    config = configparser.ConfigParser()
    config.read('config.ini')
    port = config.get('General', 'port')

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ModelService_pb2_grpc.add_ModelServiceServicer_to_server(
        ModelService(), server)
    server.add_insecure_port('[::]:' + port)
    server.start()
    print('Server is running on port ' + port)
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
