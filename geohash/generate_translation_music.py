from gradio_client import Client
import time

# Helper functions #
def timer(func):
	def wrapper(*args, **kwargs):
		start = time.time()
		rt = func(*args, **kwargs)
		end = time.time()
		print(f"{func.__name__} --> Finished:{end-start:.2f} sec")
		return rt
	return wrapper

def drawline(num = 40):
	print("-"*num)

# API functions #

@timer
def generate_translations(description:str):
	prompt = f"""
	There are three words in {description}.
	Show the definition og those words, and give an example sentence of each.
	"""
	# llama version : https://tinyllama-tinyllama-chat.hf.space/--replicas/9v2np/
	# tiny model is not stable, so lower your expectation :)
	client = Client("https://tinyllama-tinyllama-chat.hf.space/--replicas/9v2np/")
	result = client.predict(
				prompt,	# str in 'Message' Textbox component
				api_name="/chat"
				)
	drawline()
	print("Translation generated from 🦙 LLaMA:")
	print(result)
	drawline()

def main(description:str, gen_translations:bool):
	if gen_translations:
		print("Generating Translation...")
		generate_translations(description)

if __name__ == "__main__":
	# Basic Settings
	description = "('party', 'shaves', 'lance')"
	gen_translations = True
	# Run
	main(description,gen_translations)