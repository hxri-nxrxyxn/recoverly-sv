<script>
    import { Link } from "svelte-routing";
</script>

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
        </div>
    </div>
</main>

<style>
    .task {
        display: flex;
    }
    .task__adjust {
        display: flex;
    }
    .pills {
        display: flex;
        justify-content: space-between;
    }
</style>
