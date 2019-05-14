<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="20px">
      <el-form-item>
        <el-select v-model="form.appName" placeholder="选择应用">
          <el-option
            v-for="item in appName"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select
          v-model="form.hostname"
          placeholder="选择主机名"
          filterable
          allow-create
          default-first-option
        >
          <el-option
            v-for="item in applications[form.appName]"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-button type="primary" @click="onSubmit">选择</el-button>
      </el-form-item>
    </el-form>
<el-progress type="circle" :percentage=cpuP></el-progress>
<el-progress type="circle" :percentage=memP color="#8e71c7"></el-progress>
  
  </div>

</template>

<script>
export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: "success",
        draft: "gray",
        deleted: "danger"
      };
      return statusMap[status];
    }
  },
  data() {
    return {
      list: [],
      listLoading: true,
      form: {
        hostname: null,
        appName: null,
        isNewest: true
      },
      appName: [],
      applications: {},
      cpuP: 0,
      memP: 0
    };
  },
  created() {
    this.initWebSocket();
  },
  destroyed() {
    this.websock.close(); //离开路由之后断开websocket连接
  },
  methods: {
    onSubmit() {
      this.list = [];
      this.form.hostname = [this.form.hostname]
      this.websocketsend(JSON.stringify(this.form));
    },
    initWebSocket() {
      //初始化weosocket
      this.$store
        .dispatch("GetHosts")
        .then(response => {
          for (const key in response) {
            if (response.hasOwnProperty(key)) {
              let tmpvalue = [];
              for (const hostname in response[key]["hostname"]) {
                if (response[key]["hostname"].hasOwnProperty(hostname)) {
                  tmpvalue.push({ label: hostname, value: hostname });
                }
              }
              this.applications[key]=tmpvalue;
              this.appName.push({label: key, value: key})
            }
          }
          this.loading = false;
        })
        .catch(e => {
          this.loading = false;
        });
      this.listLoading = true;
      // const wsuri = "ws://localhost:8080/v1/ws";
      const wsuri = "ws://"+window.location.host+"/v1/ws";
      this.websock = new WebSocket(wsuri);
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onopen = this.websocketonopen;
      this.websock.onerror = this.websocketonerror;
      this.websock.onclose = this.websocketclose;
    },
    websocketonopen() {
      //连接建立之后执行send方法发送数据
      this.listLoading = false;
    },
    websocketonerror() {
      //连接建立失败重连
      this.initWebSocket();
    },
    websocketonmessage(e) {
      this.cpuP = JSON.parse(e.data).cpu.toFixed(2)
      this.memP = JSON.parse(e.data).mem.toFixed(2)
      //数据接收
    },
    websocketsend(Data) {
      //数据发送
      this.websock.send(Data);
    },
    websocketclose(e) {
      //关闭
      console.log("断开连接", e);
    }
  }
};
</script>
