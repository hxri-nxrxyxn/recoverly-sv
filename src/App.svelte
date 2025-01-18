<script>
  import { Router, Route, Link } from "svelte-routing";
  import Login from "./Login.svelte";
  import Register from "./Register.svelte";
  import Terms from "./Terms.svelte";
  import CollectOne from "./CollectOne.svelte";
  import CollectTwo from "./CollectTwo.svelte";
  import CollectThree from "./CollectThree.svelte";
  import Dashboard from "./Dashboard.svelte";
  import Emergency from "./Emergency.svelte";
  import Events from "./Events.svelte";
  import EventsRegister from "./EventsRegister.svelte";
  import Event from "./Event.svelte";
  import Goals from "./Goals.svelte";
  import GoalsGen from "./GoalsGen.svelte";
  import Tasks from "./Tasks.svelte";
</script>

<main>
  <Router>
    <Route path="/register" component={Register} />
    <Route path="/" component={Dashboard} />
    <Route path="/login" component={Login} />
    <Route path="/terms" component={Terms} />
    <Route path="/collect-one" component={CollectOne} />
    <Route path="/collect-two" component={CollectTwo} />
    <Route path="/collect-three" component={CollectThree} />
    <Route path="/dashboard" component={Dashboard} />
    <Route path="/emergency" component={Emergency} />
    <Route path="/events" component={Events} />
    <Route path="/events-register" component={EventsRegister} />
    <Route path="/event" component={Event} />
    <Route path="/goals" component={Goals} />
    <Route path="/goals-gen" component={GoalsGen} />
    <Route path="/tasks" component={Tasks} />
  </Router>
</main>
