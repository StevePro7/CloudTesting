import http
from flask import Flask, request, jsonify
import os

app = Flask(__name__)

#root: str = '/home/stevepro/data'
root: str = '/data'

@app.route('/')
def home():
    return "Welcome to the Flask app with persistent volume!"

@app.route('/write', methods=['POST'])
def write_file():
    data = request.json
    filename = data.get('filename', 'default.txt')
    content = data.get('content', '')

    with open(f'{root}/{filename}', 'w') as f:
        f.write(content)

    return jsonify({"message": "File written successfully!"}), http.HTTPStatus.CREATED

@app.route('/read', methods=['GET'])
def read_file():
    filename = request.args.get('filename', 'default.txt')

    if not os.path.exists(f'{root}/{filename}'):
        return jsonify({"error": "File not found!"}), http.HTTPStatus.NOT_FOUND

    with open(f'{root}/{filename}', 'r') as f:
        content = f.read()

    return jsonify({"content": content}), http.HTTPStatus.OK

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
