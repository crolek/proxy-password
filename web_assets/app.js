(function(){
	var ProxyModel = new Backbone.Epoxy.Model.extend({
		defaults: {
			username: "",
			password: ""
			/*url: "",
			port: 80,
			systemVariables: true,
			npm: false,
			bower: false,
			git: false*/
		},
		computeds: {
			results: function(){
				return this.get("username") + this.get("password");
			}
		}
	});
	
	var proxyView = new Backbone.Epoxy.View.extend({
		el: "#proxypasswordform",
		model: new ProxyModel()
	});

})()