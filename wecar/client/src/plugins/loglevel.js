import Vue from "vue"
import log from "loglevel"

log.setLevel(log.levels.DEBUG)
Vue.prototype.$log = log