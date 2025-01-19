<script>
  import { Link } from "svelte-routing";
  import { post } from "./fetch";
  import Flower from "./assets/undraw_order-flowers_81uq.png";
  import { Storage } from "@capacitor/storage";
  import NavClassic from "./NavClassic.svelte";

  let email = $state("");
  let password = $state("");

  const doPost = async () => {
    const data = {
      email: email,
      password: password,
    };
    const res = await fetch("http://192.168.183.224:8080/api/v1/login", {
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

    const { id } = await res.json();

    const setToken = async () => {
      await Storage.set({
        key: "id",
        value: id,
      });
      location.href = "/dashboard";
    };

    setToken();
  };
</script>

<main>
  <NavClassic />
  <h1>Let's beat our <span>addictions</span></h1>
  <img class="hero" src={Flower} alt="" />
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
    <Link to="/org-login">
      <button class="outlined">ORGANIZER LOGIN</button>
    </Link>
    <br />
    <button class="solid" onclick={doPost}>LOGIN</button>
    <p class="foot">
      Don't have an account? <Link to="/register"><b>Create one</b></Link>
    </p>
  </div>
</main>

<style>
</style>
