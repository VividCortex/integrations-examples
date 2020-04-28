var http = require('http')
var crypto = require('crypto')

// hookHandler listens for a POST request from VC and does something with it
//
// POST /hook
function hookHandler (req, res) {
    if (req.method == 'POST' && req.url == '/hook') {
        var data = ''
        req.on('data', function(chunk) {
            data += chunk;
        });

        req.on('end', function() {
            // Compute signature with our known Secret Token
            // https://docs.vividcortex.com/how-to-use-vividcortex/integrations/#generic-webhook
            var hash = crypto.createHash('sha1')
            hash.update(data+secretToken)
            var computedSignature = hash.digest('hex')

            var signature = req.headers['x-vividcortex-signature']

            if (signature != computedSignature) {
                console.log("signature didn't match: %s vs %s", signature, computedSignature)
                res.writeHead(401);
                res.end();
            }    

            /*

                Do Something

            */

            res.writeHead(200);
            res.end();
        })
    }
}

var app = http.createServer(hookHandler);
secretToken = 'ansecretansecretansecretansecret'

app.listen(1337);
console.log('Listening to port 1337');