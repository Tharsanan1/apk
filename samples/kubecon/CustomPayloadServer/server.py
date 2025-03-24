from flask import Flask, jsonify, request
import os

app = Flask(__name__)

@app.route('/', defaults={'path': ''}, methods=['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS'])
@app.route('/<path:path>', methods=['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS'])
def handle_request(path):
    method = request.method
    method = method.lower()
    filename = f"{method}-{path}".strip('/')  # Remove leading/trailing slashes
    filename = filename.replace('/', '-')      # Replace any remaining slashes with hyphens
    try:
        if os.path.exists(filename):
            with open(filename, 'r') as file:
                content = file.read()
                return jsonify({
                    'status': 'success',
                    'content': content
                })
        else:
            return jsonify({
                'status': 'error',
                'message': f'File {filename} not found'
            }), 404
    except Exception as e:
        return jsonify({
            'status': 'error',
            'message': str(e)
        }), 500

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5001)