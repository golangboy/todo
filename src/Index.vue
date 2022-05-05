<template>
  <div>
    <div style="display: flex">
      <el-input
        placeholder="fixBug of GoOut"
        style="margin-right: 20px; margin-left: 20px"
        v-model="title"
      ></el-input>
      <el-button type="primary" @click="addTask">Add To List</el-button>
    </div>
    <el-divider></el-divider>
    <el-empty
      description="there are nothing"
      v-if="tasks.length == 0"
    ></el-empty>
    <el-timeline reverse v-else>
      <el-timeline-item
        v-for="(item, index) in tasks"
        :key="index"
        :timestamp="item.addTime"
        size="large"
        placement="top"
      >
        <el-card>
          <div
            style="
              display: flex;
              justify-content: space-between;
              align-items: center;
            "
          >
            <div v-if="item.finishTime == ''">
              <el-tag>{{ item.title }}</el-tag>
            </div>
            <div v-else>
              Finish <el-tag type="success">{{ item.title }}</el-tag> at
              {{ item.finishTime }}
            </div>
            <div style="display: flex">
              <el-button
                size="small"
                type="success"
                @click="finishTask(index)"
                v-if="item.finishTime == ''"
                >Finish</el-button
              >
              <el-button size="small" type="danger" @click="deleteTask(index)"
                >Delete</el-button
              >
            </div>
          </div>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </div>
</template>
<script>
export default {
  name: "index",
  components: {},
  data() {
    return {
      title: "",
      tasks: [
        // {
        //   addTime: "2020-06-07 12:20:20",
        //   finishTime: "2020-06-07 12:20:20",
        //   title: "fixBug of GoOut",
        // },
        // {
        //   addTime: "2020-06-07 12:20:20",
        //   finishTime: "",
        //   title: "fixBug of GoOut",
        // },
      ],
    };
  },
  mounted: function () {
    this.read();
  },
  methods: {
    addTask() {
      this.tasks.push({
        addTime: this.dateFormatter(new Date()),
        finishTime: "",
        title: this.title,
      });
      this.save();
    },
    finishTask(index) {
      this.tasks[index].finishTime = this.dateFormatter(new Date());
      this.save();
    },
    deleteTask(index) {
      this.tasks.splice(index, 1);
      this.save();
    },
    dateFormatter(d) {
      let y = d.getFullYear();
      let m = d.getMonth() + 1;
      let day = d.getDate();
      let h = d.getHours();
      let minuns = d.getMinutes();
      let sec = d.getSeconds();
      if (m < 10) {
        m = "0" + m;
      }
      if (day < 10) {
        day = "0" + day;
      }
      if (h < 10) {
        h = "0" + h;
      }
      if (minuns < 10) {
        minuns = "0" + minuns;
      }
      if (sec < 10) {
        sec = "0" + sec;
      }
      return y + "-" + m + "-" + day + " " + h + ":" + minuns + ":" + sec;
    },
    save() {
      var xhttp = new XMLHttpRequest();
      xhttp.open("POST", "/add", false);
      xhttp.send(JSON.stringify(this.tasks));
      let obj = JSON.parse(xhttp.response);
      if (obj.err == null) {
        this.$message("OK");
      }
    },
    read() {
      var xhttp = new XMLHttpRequest();
      xhttp.open("GET", "/list", false);
      xhttp.send();
      let obj = JSON.parse(xhttp.response);
      if (obj.err == null) {
        this.$message("OK");
        this.tasks = obj.data;
      }
    },
  },
};
</script>
