from flask import Flask, jsonify, request

app = Flask(__name__)

response_data = {
    "id": "chatcmpl-Anveduo1eSWsznRXf2n7dCT4RTHu1",
    "object": "chat.completion",
    "created": 1736463059,
    "model": "gpt-4o-2024-08-06",
    "choices": [
        {
            "index": 0,
            "message": {
                "role": "assistant",
                "content": "{\"steps\":[{\"explanation\":\"We start with the equation 8x + 7 = -23. Our goal is to isolate the variable x.\",\"output\":\"8x + 7 = -23\"},{\"explanation\":\"First, we need to get rid of the constant term on the left side of the equation. We do this by subtracting 7 from both sides of the equation.\",\"output\":\"8x + 7 - 7 = -23 - 7\"},{\"explanation\":\"Simplifying both sides, we get 8x = -30.\",\"output\":\"8x = -30\"},{\"explanation\":\"The next step is to isolate x by dividing both sides of the equation by 8, since 8 is multiplied by x.\",\"output\":\"8x / 8 = -30 / 8\"},{\"explanation\":\"After dividing both sides by 8, we simplify the right side of the equation. -30/8 simplifies to -3.75 or as a fraction -15/4.\",\"output\":\"x = -3.75 or x = -15/4\"}],\"final_answer\":\"x = -3.75 or x = -15/4\"}",
                "refusal": None
            },
            "logprobs": None,
            "finish_reason": "stop"
        }
    ],
    "usage": {
        "prompt_tokens": 100,
        "completion_tokens": 238,
        "total_tokens": 338,
        "prompt_tokens_details": {
            "cached_tokens": 0,
            "audio_tokens": 0
        },
        "completion_tokens_details": {
            "reasoning_tokens": 0,
            "audio_tokens": 0,
            "accepted_prediction_tokens": 0,
            "rejected_prediction_tokens": 0
        }
    },
    "system_fingerprint": "fp_703d4ff298"
}

@app.route('/', defaults={'path': ''}, methods=["GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"])
@app.route('/<path:path>', methods=["GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"])
def catch_all(path):
    return jsonify(response_data), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000)
