var ProxyModel = Backbone.Epoxy.Model.extend({
    defaults: {
        username: "",
        password: "",
        url: "",
		port: 80,
		systemVariables: true,
		npm: false,
		bower: false,
		git: false
    },
    computeds: {
        results: function() {
            return "http://" + this.get("username") + ":" + this.get("password") + "@" + this.get("url") + ":" + this.get("port");
        }
    }
});

var view = new Backbone.Epoxy.View({
    el: "#proxypasswordform",
    model: new ProxyModel()
});