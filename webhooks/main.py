from http.server import BaseHTTPRequestHandler, HTTPServer
import hashlib

SECRET_TOKEN = "ansecretansecretansecretansecret"

""" MyWebhook listens for a POST request from VC and does something with it

POST /hook
"""
class MyWebhook(BaseHTTPRequestHandler):
    def do_POST(self):
        if self.path == '/hook':
            self.handle_webhook()
    
    def handle_webhook(self):
        content_len = int(self.headers.get('content-length'))
        post_body = self.rfile.read(content_len)

        # Compute signature with our known Secret Token
		# https://docs.vividcortex.com/how-to-use-vividcortex/integrations/#generic-webhook
        m = hashlib.sha1()
        m.update(post_body+SECRET_TOKEN.encode())
        computed_signature = m.hexdigest()

        signature = self.headers.get('X-VividCortex-Signature')

        if signature != computed_signature:
            print ("signature didn't match: %s vs %s" % (signature, computed_signature))
            self.send_response(401)
            self.end_headers()
            return True


        # 
        # 
        # Do Something
        #
        #

        self.send_response(200)
        self.end_headers()
        

if __name__ == '__main__':
    host_name = ''
    port = 1337
    httpd = HTTPServer((host_name, port), MyWebhook)
    print ("Listening on port %s" % (port))
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        pass
    httpd.server_close()

