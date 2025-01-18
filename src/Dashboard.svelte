<script>
  import { Link, useRouter } from "svelte-routing";
  import Morning from "./assets/undraw_mornings_kmib.png";
  import { Storage } from "@capacitor/storage";
  import { onMount } from "svelte";

  // navigation

  const router = useRouter();
  let hasToken = true; // Start as null to represent "loading"

  let claims;

  const checkName = async () => {
    try {
      const value = await Storage.get({ key: "id" });
      const id = value.value;
      if (!id) {
        hasToken = true;
        location.href = "/login";
      }
      const response = await fetch(
        `https://192.168.25.224:8080/api/v1/users/6`
      );

      if (response.ok) {
        console.log("response ok");
        const data = await response.json();
        console.log(data);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  checkName();
</script>

{#if hasToken === true}
  <main>
    <div class="pills">
      <div class="pill">
        <p>LEVEL</p>
      </div>
      <div class="pill">
        <p>EVENTS</p>
      </div>
      <div class="pill">
        <p>CHATROOM</p>
      </div>
    </div>
    <h1>Good morning, <span>Hari</span></h1>

    <img src={Morning} alt="" class="hero" />

    <div class="card">
      <h1>Daily Task</h1>
      <div class="tasks">
        <div class="task">
          <div class="task__text">
            <h5>No. of minutes meditated</h5>
            <p>0/15</p>
          </div>
          <div class="task__adjust">
            <button>-</button>
            <h5>0</h5>
            <button>+</button>
          </div>
        </div>
        <div class="task">
          <div class="task__text">
            <h5>No. of cigarretes smoked</h5>
            <p>0/7</p>
          </div>
          <div class="task__adjust">
            <button>-</button>
            <h5>0</h5>
            <button>+</button>
          </div>
        </div>
        <button>submit</button>
      </div>
    </div>

    <div class="achievements">
      <h2>Achivements</h2>
      <div class="achievement">
        <div class="achievement__logo">
          <i>logo</i>
        </div>
        <div class="achievement__text">
          <p>Fresh Start</p>
        </div>
      </div>
      <div class="achievement">
        <div class="achievement__logo">
          <i>logo</i>
        </div>
        <div class="achievement__text">
          <p>Fresh Start</p>
        </div>
      </div>
    </div>

    <div class="emergency">
      <Link to="/emergency">SOS</Link>
    </div>
  </main>
{:else if hasToken === false}
  <p>Redirecting to login...</p>
{:else}
  <p>Loading...</p>
{/if}

<style>
  /* Your styles */
  .task {
    display: flex;
    justify-content: space-between;
    margin-top: 1rem;
  }
  .task__adjust {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .task__adjust > * {
    margin-left: 1rem;
  }
  .task__adjust button {
    position: initial;
    font-size: 1rem;
    padding: 0.5rem;
    border: 0;
    color: var(--color-text);
    background-color: var(--color-light);
  }
  .task__adjust h5 {
    text-align: center;
  }
  .task__text p {
    font-size: 0.75rem;
  }
  .pills {
    display: flex;
    margin-bottom: 1rem;
  }
  .pill {
    background-color: var(--color-light);
    padding: 0.25rem 1rem;
    margin-right: 0.5rem;
    border-radius: 2rem;
  }
  .pill p {
    font-size: 0.5rem;
    font-weight: 700;
    color: var(--color-primary);
  }
  .achievements {
    margin-top: 2rem;
  }
  .achievement {
    display: flex;
  }
</style>
