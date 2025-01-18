<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Let it be a new <span>Beginning</span></h1>
    <div class="form__box">
        <div class="form__link">
            <label>First Name</label>
            <input type="text" placeholder="Hari" required />
        </div>
        <div class="form__link">
            <label>Last Name</label>
            <input type="password" placeholder="Narayan" required />
        </div>
        <div class="form__link">
            <label>Date of Birth</label>
            <input type="date" placeholder="DD/MM/YY" required />
        </div>
        <div class="form__link">
            <label>Gender</label>
            <input type="password" placeholder="Male" required />
        </div>
    </div>
    <div class="btn__box">
        <div class="progress">
            <div class="progressi">whoa lion!</div>
        </div>
        <Link to="/collect-two">
            <button class="solid">CONTINUE</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 30%;
    }
</style>
