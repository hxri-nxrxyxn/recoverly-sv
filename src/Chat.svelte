<script>
  import { Link } from "svelte-routing";
  import Nav from "./Nav.svelte";
  let reply = "";
  let chatsPromise = fetchChats(); // Assuming fetchChats is a function that fetches the chats

  function timeconverter(time) {
    const date = new Date(time);

    const options = {
      weekday: "short", // Short weekday (e.g., Thu)
      hour: "numeric", // Hour (numeric)
      minute: "2-digit", // Two-digit minutes
      hour12: true, // 12-hour format with AM/PM
    };
    return new Intl.DateTimeFormat("en-US", options).format(date);
  }

  async function fetchChats() {
    const response = await fetch(
      "http://localhost:8080/api/v1/chats/recovery",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    if (!response.ok) {
      const { error } = await response.json();
      alert(error);
      return;
    }

    return await response.json();
  }

  const sendReply = async () => {
    const res = await fetch("http://localhost:8080/api/v1/chat", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        message: reply,
        name: "You",
        chat_group: "recovery",
      }),
    });
    if (!res.ok) {
      const { error } = await res.json();
      alert(error);
      return;
    }

    location.reload();
  };
</script>

<main>
  <Nav />
  <h3>243 ONLINE</h3>
  <h1>Share my Story</h1>
  {#await chatsPromise}
    <p>Loading chats...</p>
  {:then chats}
    <div class="chat">
      {#each chats.data as chat}
        <div
          class="message {chat.name === 'You'
            ? 'message--out'
            : 'message--inc'}"
        >
          <p>{chat.message}</p>
          <p>
            <b>{chat.name}</b><br />
            {timeconverter(chat.sent)}
          </p>
        </div>
      {/each}
    </div>
  {:catch error}
    <p>Error loading chats: {error.message}</p>
  {/await}

  <div class="reply__box">
    <div class="reply__input">
      <input type="text" placeholder="Type something" bind:value={reply} />
    </div>
    <div class="reply__btn" on:click={sendReply}>
      <p>Send</p>
    </div>
  </div>
</main>

<style>
  .chat {
    position: relative;
  }
  .message {
    border-radius: 1rem;
    padding: 1rem;
    width: 70%;
    margin-bottom: 2rem;
    background-color: var(--color-light);
  }
  .message p:first-child {
    margin-bottom: 1rem;
  }
  .message--inc {
  }
  .message--out {
    margin-left: 60%;
    translate: -45%;
  }
  .reply__box {
    display: flex;
  }
  .reply__input {
    flex: 1;
  }
  .reply__input input {
    border-radius: 4rem;
  }
  .reply__btn {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0.5rem;
  }
</style>
