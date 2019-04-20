<template>
  <div class="dashboard-container">
    <div class="dashboard-text">使用指南</div>
    <div class="dashboard-text">1.下载客户端</div>
    <a href="https://jolyne.club" download="filebeat">下载链接</a>
    <p>或者 </p>
    <code>wget https://jolyne.club/download/filebeat</code>
    <div class="dashboard-text">2.在客户端运行</div><code>{{this.name}}</code>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      name: null
    }
  },
  created() {
    this.name = "./filebeat -f `log file` -n `app name` -u " + window.location.hostname + ":9092 -p 75e24bd1e3c189a68407f8b3c432a93f5e54564db56d6eefe4c56bdc6ef937c7"
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
