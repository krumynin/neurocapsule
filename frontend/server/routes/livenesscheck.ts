export default eventHandler((event) => {
  setResponseStatus(event, 200);

  return 'livenesscheck';
});
