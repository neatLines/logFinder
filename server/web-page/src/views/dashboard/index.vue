<template>
  <div class="dashboard-container">
    <div class="dashboard-text">功能介绍</div>
    <p>本平台为一个集中式日志收集分析平台来对日志数据进行统一的实时收集、集中存储以及简单分析。</p>
      <p>满足在业务发展，系统规模不断扩大，各种服务也越来越复杂，某些应用跨机房，跨平台部署的情况下，记录用户行为、系统运行状态的需求。可以作为业务发展的参考，也可以作为开发、测试和运维人员调查问题、诊断故障的主要依据。</p>
<p>主要满足以下需求：</p>
<p>单应用多服务器的分布式部署情况下对不同服务器内文件日志的收集分析。</p>
<p>开发测试运维人员分析日志时需要对过多的日志进行的各种过滤操作。</p>
<p>对日志做简单的分类分析。</p>
<p>便于部署，占用资源较少</p>
    <div class="dashboard-text">1.下载客户端</div>
    <el-button href=" wget https://github.com/neatLines/logFinder/raw/master/client/fileBeat" download="filebeat" type="primary">Linux可执行文件下载链接<i class="el-icon-download el-icon--right"></i></el-button>
    <el-button href=" wget https://github.com/neatLines/logFinder/raw/master/client/fileBeat" download="filebeat" type="primary">Windows可执行文件下载链接<i class="el-icon-download el-icon--right"></i></el-button>
    <el-button href=" wget https://github.com/neatLines/logFinder/raw/master/client/fileBeat" download="filebeat" type="primary">Mac OS可执行文件下载链接<i class="el-icon-download el-icon--right"></i></el-button>
<!-- 
    <a href=" wget https://github.com/neatLines/logFinder/raw/master/client/fileBeat" download="filebeat">Linux可执行文件下载链接</a>
        <a href=" wget https://github.com/neatLines/logFinder/raw/master/client/fileBeat" download="filebeat">Windows可执行文件下载链接</a>

    <a href=" wget https://github.com/neatLines/logFinder/raw/master/client/fileBeat" download="filebeat">Mac OS可执行文件下载链接</a> -->

    <p>或者 </p>
    <code>wget wget https://github.com/neatLines/logFinder/raw/master/client/fileBeat</code>
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
    this.name = "./fileBeat -u " + window.location.hostname + " -f nginx.log -wp 8080 -kp 9092 -n nginxLog -t 75e24bd1e3c189a68407f8b3c432a93f5e54564db56d6eefe4c56bdc6ef937c7"
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
