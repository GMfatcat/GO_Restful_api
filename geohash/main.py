"""
FASTAPI SERVER
"""
# Third party Import #

# Built-in Import #

# Self-import #
import API.AI.service as ai_service
import API.Query.service as query_service

# Functions #
def main():
	print("Loading",ai_service.Service(),query_service.Service())


if __name__ == "__main__":
	main()