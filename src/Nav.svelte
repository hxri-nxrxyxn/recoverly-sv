<script>
    import Star from "./assets/noun-star-923969.svg";
    import ProfilePic from "./assets/circle-user.svg";
</script>

<main>
    <div class="nav">
        <div class="nav__icon">
            <img src={Star} alt="" class="icon" />
        </div>
        <div class="nav__placeholder">hey</div>
        <div class="nav__stats">
            <h5>Hari</h5>
            <img src={ProfilePic} alt="" class="icon" />
        </div>
    </div>
</main>

<style>
    .nav {
        display: flex;
        background-color: white;
        z-index: 99;
        margin-bottom: 0.5rem;
    }
    .nav__placeholder {
        flex: 1;
        color: transparent;
    }
    .nav__icon img {
        scale: 1.2;
    }
    .nav__stats {
        display: flex;
        align-items: flex-start;
        gap: 0.5rem;
    }
</style>
