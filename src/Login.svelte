<script>
    import { Link } from "svelte-routing";
    import { post } from "./fetch";
    import Flower from "./assets/undraw_order-flowers_81uq.png";
    import { Storage } from "@capacitor/storage";

    let email = $state("");
    let password = $state("");

    const doPost = async () => {
        const data = {
            email: email,
            password: password,
        };
        // const res = await post("/api/v1/login", data);
        // const id = res.data.id;

        //  const setToken = async () => {
        //      await Storage.set({
        //          key: "id",
        //          value: id,
        //      });
        //      alert("dashboard");
        //  };

        //  setToken();
    };
</script>

<main>
    <h1>Let's beat our <span>addictions</span></h1>
    <img class="hero" src={Flower} alt="" />
    <div class="form__box">
        <div class="form__link">
            <label>Email Id</label>
            <input
                type="text"
                placeholder="Email"
                required
                bind:value={email}
            />
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
            Don't have an account?<Link to="/register"><b>Create one</b></Link>
        </p>
    </div>
</main>

<style>
</style>
