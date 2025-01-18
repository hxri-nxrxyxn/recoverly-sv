<script>
  import { Router, Route, Link } from "svelte-routing";
  import Login from "./Login.svelte";
  import Register from "./Register.svelte";
  import Terms from "./Terms.svelte";
</script>

<main>
  <Router>
    <Route path="/register" component={Register} />
    <Route path="/" component={Login} />
    <Route path="/terms" component={Terms} />
  </Router>
</main>

<style>
</style>
