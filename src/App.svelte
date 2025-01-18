<script>
  import { Router, Route, Link } from "svelte-routing";
  import Login from "./Login.svelte";
  import Register from "./Register.svelte";
</script>

<main>
  <Router>
    <Route path="/register" component={Register} />
    <Route path="/" component={Login} />
  </Router>
</main>

<style>
</style>
