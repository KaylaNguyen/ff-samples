const functions = require('@google-cloud/functions-framework');

exports.helloWorld = (req, res) => {
  throw new Error('Intended Exception')
};

functions.cloudEvent('helloCloudEvents', (cloudevent) => {
  console.log(cloudevent.id);
  console.error("intended error")
  // throw new Error('Intended Exception')
});