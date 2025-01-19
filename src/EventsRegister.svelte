<script>
  import { Link } from "svelte-routing";
  import { Storage } from "@capacitor/storage";
  import Nav from "./Nav.svelte";

  const loadevents = async () => {
    const response = await fetch("http://localhost:8080/api/v1/events", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      const { error } = await response.json();
      alert(error);
      return;
    }

    return await response.json();
  };

  const data = loadevents();
</script>

<main>
  <Nav />
  <h1>Local events <br /><span>Near You</span></h1>

  {#await data}
    <p>Loading events...</p>
  {:then events}
    {#each events.data as event}
      <div class="card">
        <div class="card__image">
          <img src={event.image || "#"} alt={event.title} />
        </div>
        <div class="card__text">
          <h1>{event.name}</h1>
          <p>{event.orginizer}</p>
          <p>{event.location}</p>
        </div>
        <button
          on:click={async () => {
            // Use the event's `id` dynamically in the API call
            console.log(event.id);
            const response = await fetch(
              `http://localhost:8080/api/v1/event/${event.id}`,
              {
                method: "GET",
                headers: {
                  "Content-Type": "application/json",
                },
              },
            );

            if (!response.ok) {
              const error = await response.json();
              console.log(error);
              return;
            }

            const data = await response.json();
            console.log(data);

            location.href = "/register/" + event.id;
          }}>REGISTER</button
        >
      </div>
    {/each}
  {:catch error}
    <p>Error loading events: {error.message}</p>
  {/await}
</main>

<style>
  .card {
    margin-bottom: 2rem;
  }
</style>
