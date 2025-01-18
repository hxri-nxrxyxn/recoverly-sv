<script>
  import { Link } from "svelte-routing";
  import NavClassic from "./NavClassic.svelte";
  import { Storage } from "@capacitor/storage";
  let password = $state("");
  let email = $state("");

  const signup = async () => {
    const data = {
      email: email,
      password: password,
    };
    const res = await fetch("http://localhost:8080/api/v1/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (!res.ok) {
      const { message } = await res.json();
      alert(message);
      return;
    }

    const newdata = await res.json();
    const id = newdata.data.id;
    const setToken = async () => {
      await Storage.set({
        key: "id",
        value: id,
      });
    };

    console.log(id);
    setToken();

    location.href = "/terms";
  };
</script>

<main>
  <NavClassic />
  <h1>Let it be a new <span>Beginning</span></h1>
  <div class="form__box">
    <div class="form__link">
      <label>Email Id</label>
      <input type="text" placeholder="Email" required bind:value={email} />
    </div>
    <div class="form__link">
      <label>Password</label>
      <input
        type="password"
        placeholder="Password"
        required
        bind:value={password}
      />
    </div>
  </div>
  <div class="btn__box">
    <button class="solid" onclick={signup}>CONTINUE</button>
    <p class="foot">Check your email for an <span>OTP</span></p>
  </div>
</main>

<style>
</style>
