"""
FASTAPI SERVER
"""
# Third party Import #
# from fastapi import FastAPI

# Built-in Import #

# Self-import #
import API.AI.service as ai_service
import API.Query.service as query_service

# # Functions #
def test():
	# module load test
	print("Loading:",ai_service.Service(),query_service.Service())

from fastapi import FastAPI

app = FastAPI()


@app.get("/")
def read_root():
	test()
	return {"Hello": "World"}

