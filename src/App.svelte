<script>
  import { Router, Route, Link } from "svelte-routing";
  import Login from "./Login.svelte";
  import Register from "./Register.svelte";
  import Terms from "./Terms.svelte";
  import CollectOne from "./CollectOne.svelte";
  import CollectTwo from "./CollectTwo.svelte";
  import CollectThree from "./CollectThree.svelte";
  import Dashboard from "./Dashboard.svelte";
</script>

<main>
  <Router>
    <Route path="/register" component={Register} />
    <Route path="/" component={Login} />
    <Route path="/terms" component={Terms} />
    <Route path="/collect-one" component={CollectOne} />
    <Route path="/collect-two" component={CollectTwo} />
    <Route path="/collect-three" component={CollectThree} />
    <Route path="/dashboard" component={Dashboard} />
  </Router>
</main>
