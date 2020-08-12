<template>
  <el-dialog
    :visible.sync="dialogVisible"
    :before-close="handleClose"
    class="add"
  >
    <p slot="title" class="add-title">Add a car</p>

    <el-form class="add-form" ref="form" :model="form" label-position="left" label-width="100px">
      <el-form-item label="Model:">
        <el-input v-model="form.model"></el-input>
      </el-form-item>
      <el-form-item label="Make:">
        <el-input v-model="form.make"></el-input>
      </el-form-item>
      <el-form-item label="Colour:">
        <el-input v-model="form.colour"></el-input>
      </el-form-item>
      <el-form-item label="Owner:">
        <el-input v-model="form.owner"></el-input>
      </el-form-item>
    </el-form>

    <span slot="footer" class="dialog-footer">
      <el-button @click="hide">Cancel</el-button>
      <el-button type="primary" @click="confirm" :loading="requesting">Confirm</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  data() {
    return {
      dialogVisible: false,
      form: {model:"", make:"", colour:"", owner:""},
      requesting: false
    };
  },
  methods: {
    reset(){
      this.form = {model:"", make:"", colour:"", owner:""}
    },
    show() {
      this.dialogVisible = true;
    },
    hide() {
      this.dialogVisible = false;
      this.reset()
    },
    confirm(){
      if(!this.form.model || !this.form.make || !this.form.colour || !this.form.owner){
        this.$alert("Please fill all boxes!", "Tips", {
          confirmButtonText: 'OK',
          callback: () => {}
        })
        return 
      }

      this.requesting = true
      this.$axios.post("/createCar", this.form)
        .then(()=> {
          this.$message({
            message: 'The car has added!.',
            type: 'success'
          })
          this.$emit("addEvent")
        })
        .then(() => {
          this.dialogVisible = false;
          this.reset()
        })
        .catch(err => {
          this.$message({
            message: err,
            type: 'error'
          })
          console.error(err)
        })
        .finally(() => {
          this.requesting = false
        })

    },
    handleClose(done) {
      this.reset()
      done()
    }
  }
};
</script>

<style lang="postcss" scoped>
.add {
  & >>> .el-dialog {
    width: 500px;
  }
  & >>> .el-dialog__header{
    padding: 5px 30px;
    border-bottom: 1px solid #C0C4CC;
  }
  &-title{
    text-align: left;
    font-weight: bold;
  }
  &-form{
    margin: auto;
    width: 350px;
  }
  .test{
  }
}
</style>