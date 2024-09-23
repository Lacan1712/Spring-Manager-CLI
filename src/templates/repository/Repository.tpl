package {{.PackageName}};

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface {{.RepositoryName}} extends JpaRepository<T, Integer> {
    Optional<T> findByEmail(String email);

    Boolean existsByEmail(String email);

    Boolean existsByCpf(String cpf);

    Optional<T> findByCpf(String cpf);

    List<T> findByNome(String nome);

    List<T> findByNomeAndSobrenome(String nome, String sobrenome);

    List<T> findAllByOrderByNomeAsc();

    Long countByNome(String nome);

    void deleteByEmail(String email);

    void deleteByCpf(String cpf);
}
