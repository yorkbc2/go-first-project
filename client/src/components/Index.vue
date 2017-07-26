<template>
  <div>
    <h1>
      Welcome to JSON website
    </h1>
    <form v-on:submit.prevent="submitForm($event)">
      <div>
        <input type="text" v-model="value" placeholder="Value">
      </div>
    </form>
    <ul>
      <li v-for="val in values">
        {{val.ID}} {{val.Value}}
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    name: "Index",

    data() {
      return {
        url: "http://localhost:8000/",
        values: [],
        value: ""
      }
    },

    methods: {

      parse(response) {
        response = JSON.parse(response.trim().replace(/[\ ]/gi, ""))

        return response
      },

      submitForm(ev) {

        this.$http.post(this.url + "post", {Value: this.value}, {emulateJSON: true})
          .then(res => {
            console.log(res)

            this.values.push({Value: res.body.Value, ID: res.body.ID})

          }, err => console.error(err))

      }

    },

    mounted() {

      this.$http.get(this.url + "api", null, {
        headers: {
          Accept: "application/json"
        }
      })
        .then(res => {

          this.values = res.body
        }, err => console.error(err))
    }
  }
</script>
