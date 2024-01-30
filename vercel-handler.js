// vercel-handler.js
module.exports = (req, res) => {
  // Redirect all requests to the Go serverless function
  const { spawn } = require('child_process');
  const goProcess = spawn('./vercel-go-function');

  goProcess.stdout.on('data', (data) => {
    console.log(`Go Process: ${data}`);
  });

  goProcess.stderr.on('data', (data) => {
    console.error(`Go Process Error: ${data}`);
  });

  res.end('Running Go reverse proxy via Vercel serverless function');
};
