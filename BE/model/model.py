import torch
import clip
from PIL import Image
from io import BytesIO


class Model:
    def __init__(self):
        self.device = "cuda" if torch.cuda.is_available() else "cpu"
        self.model, self.preprocess = clip.load("ViT-B/32", device=self.device)

    def get_embedded_image(self, data):
        data = Image.open(BytesIO(data))
        data = self.preprocess(data).unsqueeze(0).to(self.device)
        with torch.no_grad():
            features = self.model.encode_image(data)
        return features[0].tolist()

    def get_embedded_text(self, data):
        data = clip.tokenize(data).to(self.device)
        with torch.no_grad():
            features = self.model.encode_text(data)
        return features[0].tolist()
