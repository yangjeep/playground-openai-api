import openai
import os
from dotenv import load_dotenv


openai.api_key = "YOUR_API_KEY"
# get API key from top-right dropdown on OpenAI website


def main():

    read_env_file()

    openai.Engine.list()  # check we have authenticated

    embedding = openai.Embedding.create(
        input="product name", model="text-embedding-ada-002")["data"][0]["embedding"]
    print(embedding)


def read_env_file():
    # Load the .env file
    load_dotenv()

    # Get the value of OPENAI_API_KEY
    openai_api_key = os.getenv("OPENAI_API_KEY")

    if openai_api_key:
        print(f"OPENAI_API_KEY: {openai_api_key}")
        openai.api_key = openai_api_key
    else:
        print("OPENAI_API_KEY not found in .env file")


if __name__ == "__main__":
    main()
