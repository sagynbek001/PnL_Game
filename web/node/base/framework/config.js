const convict = require('convict');
const conf = convict({});

conf.loadFile('base/default-config.json');

module.exports.get = function get(propName) {
	return conf.get(propName);
};

module.exports.set = function set(propName, propValue) {
	return conf.set(propName, propValue);
}

module.exports.getPNLHost = function get() {
	return conf.get('rest_host_config.' + conf.get('app_key') + '.host');
};

module.exports.getBackendContext = function get() {
	return conf.get('rest_host_config.' + conf.get('app_key') + '.backend_context');
};
